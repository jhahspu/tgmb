let movies = [];

const getData = async () => {
  const response = await fetch('/v1/random');
  const data = await response.json();
  return data
}

const makeCard = (mv) => {
  let card = document.createElement('a')
  card.setAttribute('href', `movie/${mv.slug}`)
  card.classList.add('card')
  card.style.backgroundImage = `url(https://www.themoviedb.org/t/p/w1920_and_h800_multi_faces${mv.backdrop_path})`

  let cardDetails = document.createElement('div')
  cardDetails.classList.add('card-details')
  
  let span1 = document.createElement('span')
  span1.innerText = mv.title
  cardDetails.appendChild(span1)

  let span2 = document.createElement('span')
  span2.innerText = mv.release_date + ' | ' + mv.runtime
  cardDetails.appendChild(span2)
  
  card.appendChild(cardDetails)
  observer.observe(card)
  document.querySelector('.card-container').appendChild(card)
}

const observer = new IntersectionObserver(entries => {
  entries.forEach(entry => {
    entry.target.classList.toggle('show', entry.isIntersecting)
    if (entry.isIntersecting) observer.unobserve(entry.target)
  })
},{})

const lastCardObserver = new IntersectionObserver(entries => {
  const lastCard = entries[0]
  if (!lastCard.isIntersecting) return
  getData().then(mvs => {
    mvs.forEach(el => {
      movies.push(el)
      sessionStorage.setItem('mvs', JSON.stringify(movies))
      makeCard(el)
    })
  })
  lastCardObserver.unobserve(lastCard.target)
  lastCardObserver.observe(document.querySelector('.card:last-child'))
}, {})



// PAGE LOADED
document.addEventListener('DOMContentLoaded', () => {
  if (sessionStorage.getItem('mvs') && sessionStorage.getItem('mvs').length > 0) {
    movies = JSON.parse(sessionStorage.getItem('mvs'))
    movies.forEach(el => {
      makeCard(el)
    })
    lastCardObserver.observe(document.querySelector('.card:last-child'))
  } else {
    getData().then(mvs => {
      mvs.forEach(el => {
        movies.push(el)
        sessionStorage.setItem('mvs', JSON.stringify(movies))
        makeCard(el)
      })
    }).then(() => lastCardObserver.observe(document.querySelector('.card:last-child')))
  }
  
})
