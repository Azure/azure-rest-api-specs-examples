
import com.azure.resourcemanager.networkanalytics.models.ContainerSaS;
import java.time.OffsetDateTime;

/**
 * Samples for DataTypes GenerateStorageContainerSasToken.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataTypes_GenerateStorageContainerSasToken_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataTypes_GenerateStorageContainerSasToken_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataTypesGenerateStorageContainerSasTokenMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataTypes().generateStorageContainerSasTokenWithResponse("aoiresourceGroupName", "dataproduct01",
            "datatypename",
            new ContainerSaS().withStartTimestamp(OffsetDateTime.parse("2023-08-24T05:34:58.039Z"))
                .withExpiryTimestamp(OffsetDateTime.parse("2023-08-24T05:34:58.039Z")).withIpAddress("1.1.1.1"),
            com.azure.core.util.Context.NONE);
    }
}
