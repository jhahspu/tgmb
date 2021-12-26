const t = document.querySelector('.trailers')
const trailers = t.dataset.trailers.split(' ')

const makeIframe = (a) => {
  let d = document.createElement('div')
  let i = document.createElement('iframe')
  i.setAttribute('src', `https://www.youtube.com/embed/${a}?&autoplay=0&loop=1&rel=0&showinfo=0&color=white&iv_load_policy=3`)

  d.appendChild(i)
  t.appendChild(d)
}

trailers.forEach(s => {
  if (s.length > 0) {
    makeIframe(s)
  }
});
