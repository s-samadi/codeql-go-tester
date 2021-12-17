import go

from Function println, DataFlow::CallNode call
where
  println.hasQualifiedName("fmt", "Println") and
  call = println.getACall()
select call