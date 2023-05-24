import com.azure.resourcemanager.databox.models.ShipmentPickUpRequest;
import java.time.OffsetDateTime;

/** Samples for Jobs BookShipmentPickUp. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/BookShipmentPickupPost.json
     */
    /**
     * Sample code: BookShipmentPickupPost.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void bookShipmentPickupPost(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .bookShipmentPickUpWithResponse(
                "YourResourceGroupName",
                "TestJobName1",
                new ShipmentPickUpRequest()
                    .withStartTime(OffsetDateTime.parse("2019-09-20T18:30:00Z"))
                    .withEndTime(OffsetDateTime.parse("2019-09-22T18:30:00Z"))
                    .withShipmentLocation("Front desk"),
                com.azure.core.util.Context.NONE);
    }
}
