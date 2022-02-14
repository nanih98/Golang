import requests
import requests

def main():
    r = requests.post('https://people.zoho.com/people/api/attendance?authtoken=1000.d93d1419485f8283f8c3a69346288685.61257eaa2fbd9677eb3cfa438d45185d&dateFormat=dd/MM/yyyy+HH:mm:ss&checkIn=11/02/2022+11:59:25&mapId=dromero@edrans.com')
    print(r.text)
    print(r.status_code)


main()