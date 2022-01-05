"use strict";

let editor = new FroalaEditor('#show', {
    // https://froala.com/wysiwyg-editor/docs/options
    // Char Counter
    charCounterCount: false,

    // General
    attribution: false,
    fullPage: true,
    iframe: true,
    placeholderText: "",
    spellcheck: false,
    toolbarButtons: [],
    toolbarButtonsSM: [],
    toolbarButtonsXS: [],

    // Image
    imageDefaultAlign: 'left',
    imageInsertButtons: ['imageUpload', 'imageManager'],

    // Link
    linkAlwaysBlank: true,
    linkInsertButtons: [],
    linkText: true,

    // Quick Insert
    quickInsertEnabled: false,
});


$("#save").click(function () {
    $("#content").val(editor.html.get());
});

