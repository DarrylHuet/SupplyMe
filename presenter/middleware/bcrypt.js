
bcrypt.hash(password, rounds, (err, hash) => {
    if (err) {
      console.error(err)
      return
    }
    console.log(hash)
  })

  bcrypt.compare(password, hash, (err, res) => {
    if (err) {
      console.error(err)
      return
    }
    console.log(res) //true or false
  })