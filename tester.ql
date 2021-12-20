import go   
import Mux

from Method get, RequestVars vars, DataFlow::CallNode call
where
  get.hasQualifiedName("net/http", "Header", "Get") and
  call = get.getACall()
select call, vars
