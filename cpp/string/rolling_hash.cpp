#include<string>
#include<vector>
#include<algorithm>
#include<iostream>

using namespace std;

typedef unsigned long long ull;

struct RollingHash{
  static const ull p = 100000007;
  string s;
  int n;
  ull *pow, *phash;

  RollingHash(const string &s) : s(s), n(s.size()) {
    pow = (ull *)malloc((n + 1) * sizeof(ull));
    phash = (ull *)malloc((n + 1) * sizeof(ull));
    pow[0] = 1;
    phash[0] = 0;
    for(int i = 0; i < n; i++){
      phash[i+1] = s[i] + phash[i] * p;
      pow[i+1] = pow[i] * p;
    }
  }

  bool operator()(const int i, const int j) { 
    int k = lcp(i, j);
    if(n <= i + k){
      return true;
    }
    if (n <= j + k){
      return false;
    }
    return s[i+k] <= s[j+k];
  }

  static ull hash(const string &t) {
    ull h=0;
    for(int i = 0; i < t.size(); i++){
      h = t[i] + h * p;
    }
    return h;
  }

  inline ull hash(const int& b, const int& e) const {
    return phash[e] - phash[b] * pow[e-b];
  }

  inline int find(string t) {
    if(s.size() < t.size()){
      return 0;
    }
    int w = t.size(), count = 0;
    ull h=hash(t);

    for(int i = 0; i < n-w+1; i++){
      count += (hash(i, i+w) == h);
    }
    return count;
  }

  inline int lcp(const int& i, const int& j) const {
    int l = 0, r = n - max(i,j) + 1;
    while(l + 1 < r) {
      const int m = (l + r) >> 1;
      (hash(i,i+m) == hash(j,j+m) ? l : r) = m;
    }
    return l;
  }
};

vector<int> suffixArray(const RollingHash &rh){
  int n = rh.n + 1, sa[n];
  for(int i = 0; i < n; i++){
    sa[i] = i;
  }
  sort(sa, sa + n, rh);
  vector<int>res(n);
  for(int i = 0; i < n; i++){
    res[i] = sa[i];
  }
  return res;
}

bool contain(const string &S, const vector<int> &sa, const string &T){
  int a = 0, b = S.size();
  while(1 < b - a){
    const int c = (a + b)>>1;
    (S.compare(sa[c],T.length(), T) < 0 ? a : b ) = c;
  }
  return S.compare(sa[b], T.length(), T) == 0;
}

//POJ 2217: http://poj.org/problem?id=2217
int main(void){
  ios::sync_with_stdio(false);

  int n;
  string s,t;
  cin >> n;
  cin.ignore();
  while(n--){
    getline(cin,s);
    getline(cin,t);
    int sl=s.size(),ans=0;
    RollingHash rh=RollingHash(s+"\1"+t);
    vector<int>sa=suffixArray(rh);

    for(int i=0;i+1<rh.n;i++){
      if((sa[i]<sl)!=(sa[i+1]<sl))ans=max(ans,rh.lcp(sa[i],sa[i+1]));
    }
    
    cout << "Nejdelsi spolecny retezec ma delku " << ans << ".\n";
  }
  return 0;
}