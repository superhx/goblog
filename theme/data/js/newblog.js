/**
 * Created by guotianhao on 15/5/25.
 */
window.addEventListener("load", function () {
    var markDownEl = document.querySelector(".markdown");
    new MediumEditor(document.querySelector(".editor"), {
        buttons: ["bold", "italic", "underline", "orderedlist", "unorderedlist", "anchor", "header1", "header2", "quote"]
        , extensions: {
            markdown: new MeMarkdown(function (md) {
                markDownEl.innerText = md;
            })
        }
    });
});

