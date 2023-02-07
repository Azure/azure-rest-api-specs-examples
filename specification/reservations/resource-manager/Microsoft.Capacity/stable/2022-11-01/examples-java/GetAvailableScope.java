import com.azure.resourcemanager.reservations.models.AvailableScopeRequest;
import com.azure.resourcemanager.reservations.models.AvailableScopeRequestProperties;
import java.util.Arrays;

/** Samples for Reservation AvailableScopes. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetAvailableScope.json
     */
    /**
     * Sample code: AvailableScopes.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void availableScopes(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .reservations()
            .availableScopes(
                "276e7ae4-84d0-4da6-ab4b-d6b94f3557da",
                "356e7ae4-84d0-4da6-ab4b-d6b94f3557da",
                new AvailableScopeRequest()
                    .withProperties(
                        new AvailableScopeRequestProperties()
                            .withScopes(Arrays.asList("/subscriptions/efc7c997-7700-4a74-b731-55aec16c15e9"))),
                com.azure.core.util.Context.NONE);
    }
}
