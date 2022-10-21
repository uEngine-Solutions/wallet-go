package wallet.domain;

import wallet.domain.*;
import wallet.infra.AbstractEvent;
import lombok.*;
import java.util.*;
@Data
@ToString
public class InsufficientPointsOccurred extends AbstractEvent {

    private Long id;
}


