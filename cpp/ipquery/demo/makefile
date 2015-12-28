CXX=/home/haoyankai/software/gcc/bin/g++

GEN_CPP_DIR := ../../../gen/gen-cpp
GEN_SRC := $(GEN_CPP_DIR)/ipquery.cpp $(GEN_CPP_DIR)/ipquery_types.cpp $(GEN_CPP_DIR)/ipquery_constants.cpp ../IpqueryClient.cpp IpqueryTest.cpp 
GEN_OBJ := $(patsubst %.cpp,%.o, $(GEN_SRC))

THRIFT_DIR := /home/haoyankai/software/thrift/include
BOOST_DIR := /home/haoyankai/software/boost/include 
THRIFT_LIB_DIR := /home/haoyankai/software/thrift/lib

COMMON_LINKER := -lthrift

INC := -I$(THRIFT_DIR) -I$(BOOST_DIR) -I$(GEN_CPP_DIR)
 
.PHONY: all clean
 
all: client
 
%.o: %.cpp
	$(CXX) -Wall $(INC) -c $< -o $@
 
client: $(GEN_OBJ) IpqueryTest.o
	$(CXX) $^ -L$(THRIFT_LIB_DIR) $(COMMON_LINKER) -o $@

clean:
	$(RM) ../../../gen/gen-cpp/*.o ../*.o *.o client 
