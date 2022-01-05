package star

import (
	"meta-go/model"
	"meta-go/router/view"
	"meta-go/service/starservice"

	"github.com/gofiber/fiber/v2"
)

type handler struct{}

var (
	defaultOffset       = 0
	defaultLimit        = 10
	minPage       int64 = 1
	maxPage       int64 = 10

	starHandler handler
)

func genPageRange(startPage, endPage int64) []int64 {
	pageRange := []int64{}
	for endPage >= startPage {
		pageRange = append(pageRange, startPage)
		startPage += 1
	}
	return pageRange
}

func (handler) list(isMy bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &StarListRequest{}
		if err := view.ValidateRequestQuery(c, req); err != nil {
			return c.Render("star/list", view.RespWithWarn(c, err.Error()))
		}

		var listUrl string
		switch c.Path() {
		case "/star/index":
			req.Offset = defaultOffset
			req.Limit = defaultLimit
			listUrl = "/star/list"
		default:
			listUrl = c.Path()
		}

		if req.Limit == 0 {
			req.Offset = defaultOffset
			req.Limit = defaultLimit
		}

		var authorID uint
		if isMy {
			authorID = c.UserContext().Value("user").(model.User).ID
		}

		stars, count, err := starservice.GetStarList(authorID, req.Offset, req.Limit)
		if err != nil {
			c.Render("star/list", view.RespWithWarn(c, err.Error()))
		}

		// 分页
		currPage := int64(req.Offset/req.Limit + 1)
		totalPage := (count + int64(req.Limit) - 1) / int64(req.Limit)
		if currPage > totalPage {
			currPage = totalPage
		}

		var pageRange []int64
		if totalPage > 0 && totalPage <= maxPage {
			pageRange = genPageRange(minPage, totalPage)
		} else if totalPage > maxPage {
			if currPage < 7 {
				pageRange = genPageRange(minPage, maxPage)
			} else {
				if currPage+4 > totalPage {
					pageRange = genPageRange(totalPage-9, totalPage)
				} else {
					pageRange = genPageRange(currPage-5, currPage+4)
				}
			}
		}

		return c.Render("star/list", view.Resp(c, fiber.Map{
			"stars":     stars,
			"count":     count,
			"limit":     req.Limit,
			"offset":    req.Offset,
			"listUrl":   listUrl,
			"currPage":  currPage,
			"totalPage": totalPage,
			"pageRange": pageRange,
		}))
	}
}

func (handler) detail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &StarIDRequest{}
		if err := view.ValidateRequestQuery(c, req); err != nil {
			return c.Render("star/detail", view.RespWithWarn(c, err.Error()))
		}

		star, err := starservice.GetStarDetail(0, req.ID)
		if err != nil {
			return c.Render("star/detail", view.RespWithWarn(c, err.Error()))
		}
		return c.Render("star/detail", view.Resp(c, fiber.Map{"star": star}))
	}
}

func (handler) createStar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "GET" {
			return c.Render("star/add", view.Resp(c))
		}
		req := &CreateStarRequest{}
		if err := view.ValidateRequestBody(c, req); err != nil {
			return c.Render("star/add", view.RespWithWarn(c, err.Error()))
		}
		user := c.UserContext().Value("user").(model.User)
		err := starservice.AddStar(user.ID, req.Title, req.Content)
		if err != nil {
			return c.Render("star/add", view.RespWithWarn(c, err.Error()))
		}
		return c.Redirect("/star/my/list")
	}
}

func (handler) updateStar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "GET" {
			req := &StarIDRequest{}
			if err := view.ValidateRequestQuery(c, req); err != nil {
				return c.Render("star/detail", view.RespWithWarn(c, err.Error()))
			}
			user := c.UserContext().Value("user").(model.User)
			star, err := starservice.GetStarDetail(user.ID, req.ID)
			if err != nil {
				return c.Render("star/detail", view.RespWithWarn(c, err.Error()))
			}
			return c.Render("star/edit", view.Resp(c, fiber.Map{"star": star}))
		}
		// 更新
		req := &UpdateStarRequest{}
		if err := view.ValidateRequestBody(c, req); err != nil {
			return c.Render("star/edit", view.RespWithWarn(c, err.Error(), fiber.Map{"star": req}))
		}
		user := c.UserContext().Value("user").(model.User)
		err := starservice.UpdateStar(user.ID, req.ID, req.Title, req.Content)
		if err != nil {
			return c.Render("star/edit", view.RespWithWarn(c, err.Error(), fiber.Map{"star": req}))
		}
		return c.Render("star/edit", view.RespWithInfo(c, "修改成功", fiber.Map{"star": req}))
	}
}
