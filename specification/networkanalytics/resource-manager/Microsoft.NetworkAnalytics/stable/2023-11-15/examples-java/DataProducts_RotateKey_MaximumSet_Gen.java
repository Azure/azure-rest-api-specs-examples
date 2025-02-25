
import com.azure.resourcemanager.networkanalytics.models.KeyVaultInfo;

/**
 * Samples for DataProducts RotateKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProducts_RotateKey_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataProducts_RotateKey_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void
        dataProductsRotateKeyMaximumSetGen(com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataProducts().rotateKeyWithResponse("aoiresourceGroupName", "dataproduct01",
            new KeyVaultInfo().withKeyVaultUrl("fakeTokenPlaceholder"), com.azure.core.util.Context.NONE);
    }
}
