window.addEventListener('scroll', () => {
    document.querySelector('nav').classList.toggle
    ('window-scroll', window.scrollY > 0)
})

const url = "https://api.github.com/repos/Shreejan-35/WOW";
fetch(url)
.then(data=>{console.log(data.json); return data.json()})
.then(res=>{console.log(res)})