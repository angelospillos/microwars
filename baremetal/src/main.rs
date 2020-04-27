use actix_web::{App, HttpResponse, HttpServer, Responder, get};
use bigint::BigUint;
use num_traits::{Zero, One};
use tramp::{tramp, Rec};
use uuid::Uuid;

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(status)
            .service(work)
            .service(test)
            .service(combat)
            .service(jab)
            .service(cross)
            .service(hook)
            .service(uppercut)
    })
        .bind("127.0.0.1:5000")?
        .run()
        .await
}

#[get("/status")]
async fn status() -> impl Responder {
    let response =  "{ status: \"ok\"}";
    HttpResponse::Ok().body(response)
}

#[get("/work")]
async fn work() -> impl Responder {
    let response =  format!(r#"{{ "uuid" : {uuid}, "fib":{fib} }}"#, uuid = Uuid::new_v4(), fib = fib(16));
    HttpResponse::Ok().body(response)
}

#[get("/test")]
async fn test() -> impl Responder {
    let response =  format!(r#"{{ "uuid" : {uuid}, "fib":{fib} }}"#, uuid = Uuid::new_v4(), fib = fib(16));
    HttpResponse::Ok().body(response)
}

#[get("/combat")]
async fn combat() -> impl Responder {
    let response =  format!(r#"{{ "uuid" : {uuid}, "fib":{fib} }}"#, uuid = Uuid::new_v4(), fib = fib(16));
    HttpResponse::Ok().body(response)
}

#[get("/jab")]
async fn jab() -> impl Responder {
    let response =  format!(r#"{{ "uuid" : {uuid}, "fib":{fib} }}"#, uuid = Uuid::new_v4(), fib = fib(2));
    HttpResponse::Ok().body(response)
}

#[get("/cross")]
async fn cross() -> impl Responder {
    let response =  format!(r#"{{ "uuid" : {uuid}, "fib":{fib} }}"#, uuid = Uuid::new_v4(), fib = fib(4));
    HttpResponse::Ok().body(response)
}

#[get("/hook")]
async fn hook() -> impl Responder {
    let response =  format!(r#"{{ "uuid" : {uuid}, "fib":{fib} }}"#, uuid = Uuid::new_v4(), fib = fib(8));
    HttpResponse::Ok().body(response)
}

#[get("/uppercut")]
async fn uppercut() -> impl Responder {
    let response =  format!(r#"{{ "uuid" : {uuid}, "fib":{fib} }}"#, uuid = Uuid::new_v4(), fib = fib(16));
    HttpResponse::Ok().body(response)
}

extern crate num_bigint as bigint;
extern crate num_traits;
#[macro_use] extern crate tramp;

fn fib(n: i32) -> BigUint {
  tramp(do_fib(n, Zero::zero(), One::one()))
}

fn do_fib(n: i32, acc: BigUint, curr: BigUint) -> Rec<BigUint> {
  if n <= 0 {
    rec_ret!(acc)
  }else{
    let new = &acc + curr;
    let nn = n - 1;
    rec_call!(do_fib(nn, new, acc))
  }
}