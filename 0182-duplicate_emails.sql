select email as "Email"
from Person p
group by email
having count(email) > 1