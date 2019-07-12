CXX			:= g++
CXXFLAGS	:= -Wall -std=c++17 -O2 -fPIC

TARGET		:= moridemo
PROJ_ROOT	:= $(realpath .)
INCLUDES	:= 
SRCS		:= main.cpp led.cpp
LIBS		:= -lczmq -lzmq
OBJS		:= $(SRCS:.cpp=.o)

### Rules ######################################################################

.PHONY: all
all: $(TARGET)

.PHONY: make
make: $(TARGET)

$(TARGET): $(OBJS)
	$(CXX) -o $@ $(OBJS) $(LIBS)

.PHONY: clean
clean:
	$(RM) $(TARGET) *.o

.cpp.o:
	$(CXX) -c $(CXXFLAGS) $(INCLUDES) $<
