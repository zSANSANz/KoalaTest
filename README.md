# KoalaTest

1. Logic Test
o Please create business logic for the fibonaci number sequence (0-1000).
o Please create business logic for prima number (0-1000).
o Please create business logic for check is palindrome.
2. Database Test
1. From above ERD please write query
a. Display Customer List including calculating the total order.
b. Show Product List including calculating the number of orders sorted
by the most in the order.
c. Display the sort payment method data most frequently used by
customers.
3. Rest Test.
• From Above ERD please create Rest full API.
1. Create register API(Include Generate password).
• Acceptance
o Phone number and email is unique.
o Customer name,email,phone number,dob,sex,created_date
is mandatory.
o Password generated using SHA256 etc mix with salt
key(dynamic).
o
• Negative case
2. Create get token api.
• Acceptance
o Phone_number_or_email and password is mandatory.
o Passed validation from (phone_number_or_email ) and
password.
o Must return token with access & refresh type
3. Create refresh token api.
• Acceptance.
o Must return token with access & refresh type
4. Create order api.
• Acceptance
o Passed validation from bearer auth.
o token is only one use.
o order number generated with format PO-123/IX/2020 (IX is
current month)(2020 is current year),(123 reset per month).
o order detail can be more than one
