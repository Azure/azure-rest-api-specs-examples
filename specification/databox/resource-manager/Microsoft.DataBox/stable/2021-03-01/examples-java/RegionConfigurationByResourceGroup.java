import com.azure.resourcemanager.databox.models.DataBoxScheduleAvailabilityRequest;
import com.azure.resourcemanager.databox.models.RegionConfigurationRequest;

/** Samples for Service RegionConfigurationByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/RegionConfigurationByResourceGroup.json
     */
    /**
     * Sample code: RegionConfigurationByResourceGroup.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void regionConfigurationByResourceGroup(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .services()
            .regionConfigurationByResourceGroupWithResponse(
                "SdkRg4981",
                "westus",
                new RegionConfigurationRequest()
                    .withScheduleAvailabilityRequest(
                        new DataBoxScheduleAvailabilityRequest().withStorageLocation("westus")),
                com.azure.core.util.Context.NONE);
    }
}
