"use strict";

$(document).ready(function () {
    let toolbarButtons = {
        'moreText': {
            'buttons': ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', 'textColor', 'backgroundColor', 'inlineClass', 'inlineStyle', 'clearFormatting']
        },
        'moreParagraph': {
            'buttons': ['alignLeft', 'alignCenter', 'formatOLSimple', 'alignRight', 'alignJustify', 'formatOL', 'formatUL', 'paragraphFormat', 'paragraphStyle', 'lineHeight', 'outdent', 'indent', 'quote']
        },
        'moreRich': {
            'buttons': ['insertImage', 'insertTable', 'emoticons', 'insertHR', 'specialCharacters', 'markdown'],
            'buttonsVisible': 2
        },
        'moreMisc': {
            'buttons': ['undo', 'redo', 'fullscreen', 'print', 'selectAll', 'html', 'help'],
            'align': 'right',
            'buttonsVisible': 2
        }
    }

    let toolbarButtonsXS = {
        'moreText': {
            'buttons': ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', 'textColor', 'backgroundColor', 'inlineClass', 'inlineStyle', 'clearFormatting'],
            'buttonsVisible': 0
        },
        'moreParagraph': {
            'buttons': ['alignLeft', 'alignCenter', 'formatOLSimple', 'alignRight', 'alignJustify', 'formatOL', 'formatUL', 'paragraphFormat', 'paragraphStyle', 'lineHeight', 'outdent', 'indent', 'quote'],
            'buttonsVisible': 0
        },
        'moreRich': {
            'buttons': ['insertImage', 'insertTable', 'emoticons', 'insertHR', 'specialCharacters', 'markdown'],
            'buttonsVisible': 0
        },
        'moreMisc': {
            'buttons': ['undo', 'redo', 'fullscreen', 'print', 'selectAll', 'html', 'help'],
            'align': 'right',
            'buttonsVisible': 2
        }
    }
    let editor = new FroalaEditor('#editor', {
        // https://froala.com/wysiwyg-editor/docs/options
        // Char Counter
        charCounterCount: false,

        // General
        attribution: false,
        // fullPage: true,
        height: 400,
        // iframe: true,
        spellcheck: false,
        tabSpaces: 4,
        toolbarButtons: toolbarButtons,
        toolbarButtonsSM: toolbarButtons,
        toolbarButtonsXS: toolbarButtonsXS,

        // Image
        imageDefaultAlign: 'left',
        imageInsertButtons: ['imageUpload', 'imageManager'],

        // Line Breaker
        lineBreakerTags: ['table', 'hr', 'form', 'dl', 'span.fr-video', '.fr-embedly', '.fr-img-caption', 'pre'],

        // Link
        linkAlwaysBlank: true,
        linkInsertButtons: [],
        linkText: true,

        // Paragraph Format
        paragraphFormat:{
            N: 'Normal',
            H1: 'Heading 1',
            H2: 'Heading 2',
            H3: 'Heading 3',
            H4: 'Heading 4',
            PRE: 'Pre'
        },

        // Quick Insert
        quickInsertEnabled: false,
    });

    $("#save").click(function () {
        $("#content").text(editor.html.get());
    });
});