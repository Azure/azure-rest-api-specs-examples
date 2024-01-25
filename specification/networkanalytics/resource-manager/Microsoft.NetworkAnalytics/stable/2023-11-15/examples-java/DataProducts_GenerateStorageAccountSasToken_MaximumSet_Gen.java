
import com.azure.resourcemanager.networkanalytics.models.AccountSas;
import java.time.OffsetDateTime;

/**
 * Samples for DataProducts GenerateStorageAccountSasToken.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProducts_GenerateStorageAccountSasToken_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataProducts_GenerateStorageAccountSasToken_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataProductsGenerateStorageAccountSasTokenMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataProducts().generateStorageAccountSasTokenWithResponse("aoiresourceGroupName", "dataproduct01",
            new AccountSas().withStartTimestamp(OffsetDateTime.parse("2023-08-24T05:34:58.151Z"))
                .withExpiryTimestamp(OffsetDateTime.parse("2023-08-24T05:34:58.151Z")).withIpAddress("1.1.1.1"),
            com.azure.core.util.Context.NONE);
    }
}
