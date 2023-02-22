from bs4 import BeautifulSoup
import requests
import time

print('Put the base skill')
base_skill = input('>')
print(f'Finding {base_skill} jobs')

def find_jobs():
    html_text = requests.get('https://www.timesjobs.com/candidate/job-search.html?searchType=Home_Search&from=submit&asKey=OFF&txtKeywords=&cboPresFuncArea=35').text
    soup = BeautifulSoup(html_text, 'lxml')
    jobs = soup.find_all('li', class_='clearfix job-bx wht-shd-bx')
    for i, job in enumerate(jobs):
        publish_date = job.find('span', class_='sim-posted').text
        if 'today' in publish_date:
            skills = job.find('span', class_='srp-skills').text.replace(' ', '')
            if base_skill in skills:
                company = job.find('h3', class_='joblist-comp-name').text.replace(' ', '')
                more_info = job.header.h2.a['href']
                with open(f'./{base_skill}_{i}.txt', 'w') as f:
                    f.write(f'Company: {company.strip()} \n')
                    f.write(f'Skills: {skills.strip()} \n')
                    f.write(f'MoreInfo: {more_info}')
                print(f'File {i} Saved!')

if __name__ == '__main__':
    while True:
        find_jobs()
        wait = 1
        print(f'Waiting {wait} minutes ...')
        time.sleep(wait * 60)