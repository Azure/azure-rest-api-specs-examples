import com.azure.resourcemanager.databox.models.DataBoxScheduleAvailabilityRequest;
import com.azure.resourcemanager.databox.models.RegionConfigurationRequest;

/** Samples for Service RegionConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/RegionConfiguration.json
     */
    /**
     * Sample code: RegionConfiguration.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void regionConfiguration(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .services()
            .regionConfigurationWithResponse(
                "westus",
                new RegionConfigurationRequest()
                    .withScheduleAvailabilityRequest(
                        new DataBoxScheduleAvailabilityRequest().withStorageLocation("westus")),
                com.azure.core.util.Context.NONE);
    }
}
