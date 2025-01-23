#include<bits/stdc++.h>
using namespace std;
/* single responsibility (a class should have only one reason to change 1 responsibility , to should a single work)*/ 
class GenerateSlip{
    Public:
    int numberOfItems;
    int PriceOfItems[];
    int QuantityOfItems[];

    int CalculateTotalPrice(){
        int totalPrice =0;
        for(int i=0;i<numberOfItems;i++){
            totalPrice += (PriceOfItems[i]*QuantityOfItems[i]);
        }
        return totalPrice;
    }

    void saveToDatabase(){
        //api call to  save this calculation into database
    }

    void PrinttheSlip(){
        //printer 
    }

};
//this is not following our single responsibility principle,  becasuse 
// it is dependent on many works or things to change


class Invoice{
    int numberOfItems;
    int PriceOfItem[];
    int QuantityOfItem[];
}

class CalculateTotalPriceInvoice{
    public:
    Invoice invoice;
    CalculateTotalPriceInvoice(Invoice invoice){
        this->invoice  = invoice
    }

    int totalPrice(){
        int totalPrice =0;
        for(int i=0;i<invoice.numberOfItems;i++){
            totalPrice += (invoice.PriceOfItem[i]*invoice.QuantityOfItem[i])
        } 
        return totalPrice
    }
};

class SaveToInvoice{
    public:
    Invoice invoice;
    saveToInvoice(Invoice invoice){
        this->invoice  = invoice
    }

    void saveToInvoice(){
       
    }
};

int main(){

}