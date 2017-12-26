from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
import SocketServer

class S(BaseHTTPRequestHandler):
    def _set_headers(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()

    def do_GET(self):

	print(self.command + ' '+ self.path + ' ' + self.request_version)
        self._set_headers()

    def do_HEAD(self):
        self._set_headers()
        
    def do_POST(self):
        # Doesn't do anything with posted data
	content_len = int(self.headers.getheader('content-length', 0))
	post_body = self.rfile.read(content_len)
	print (self.path + ' ' +  post_body)
        self._set_headers()

    def do_PUT(self):
        # Doesn't do anything with posted data
        self._set_headers()


    def do_DELETE(self):
        # Doesn't do anything with posted data
        self._set_headers()

        
def run(server_class=HTTPServer, handler_class=S, port=8080):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    print 'Starting httpd...'
    httpd.serve_forever()

if __name__ == "__main__":
    from sys import argv

    if len(argv) == 2:
        run(port=int(argv[1]))
    else:
        run()
