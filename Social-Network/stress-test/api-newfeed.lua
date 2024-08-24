local baseUrl = "http://127.0.0.1:8080"  -- Replace with your actual API endpoint

-- Định nghĩa hàm để tạo request
request = function()
    local usernames = {"username1", "username2", "username3", "username4"}
    local randomIndex = math.random(1, #usernames)
    local username = usernames[randomIndex]
  
    local headers = "session=session-" .. username
    wrk.headers["Cookie"] = headers
    path = "/v1/newsfeeds" 


   return wrk.format("GET", baseUrl .. path)

  end
  
  -- Định nghĩa hàm để xử lý response
  response = function(status, headers, body)
    if status == 200 then
      -- Xử lý response thành công
      print("Response OK:", status)
      -- Thêm xử lý response nếu cần
    else
      -- Xử lý response lỗi
      print("Response Error:", status)
      -- Thêm xử lý response lỗi nếu cần
    end
  end
  
  -- Chạy test
  wrk.method = "GET" -- Thiết lập phương thức GET
--   wrk.headers["Cookie"] = "session=session-username1" -- Thiết lập cookie
  
  