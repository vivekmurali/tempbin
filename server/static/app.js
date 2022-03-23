const article = document.getElementById("article");

//console.log(article.innerText);

article.addEventListener("click", () => {
  console.log("clicked");
  article.setAttribute("data-tooltip", "Copied!");
  navigator.clipboard.writeText(article.innerHTML)
});
