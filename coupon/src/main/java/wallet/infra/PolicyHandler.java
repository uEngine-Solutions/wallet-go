package wallet.infra;

import javax.naming.NameParser;

import javax.naming.NameParser;
import javax.transaction.Transactional;

import wallet.config.kafka.KafkaProcessor;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.stream.annotation.StreamListener;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Service;
import wallet.domain.*;


@Service
@Transactional
public class PolicyHandler{
    @Autowired CouponRepository couponRepository;
    
    @StreamListener(KafkaProcessor.INPUT)
    public void whatever(@Payload String eventString){}

    @StreamListener(value=KafkaProcessor.INPUT, condition="headers['type']=='InsufficientPointsOccurred'")
    public void wheneverInsufficientPointsOccurred_Cancel(@Payload InsufficientPointsOccurred insufficientPointsOccurred){

        InsufficientPointsOccurred event = insufficientPointsOccurred;
        System.out.println("\n\n##### listener Cancel : " + insufficientPointsOccurred + "\n\n");


        

        // Sample Logic //
        Coupon.cancel(event);
        

        

    }

}


