
import com.azure.resourcemanager.providerhub.models.SkuCost;
import com.azure.resourcemanager.providerhub.models.SkuResourceProperties;
import com.azure.resourcemanager.providerhub.models.SkuSetting;
import java.util.Arrays;

/**
 * Samples for Skus CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_CreateOrUpdate.
     * json
     */
    /**
     * Sample code: Skus_CreateOrUpdate.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().define("testSku").withExistingResourcetypeRegistration("Microsoft.Contoso", "testResourceType")
            .withProperties(new SkuResourceProperties().withSkuSettings(
                Arrays.asList(new SkuSetting().withName("freeSku").withTier("Tier1").withKind("Standard"),
                    new SkuSetting().withName("premiumSku").withTier("Tier2").withKind("Premium")
                        .withCosts(Arrays.asList(new SkuCost().withMeterId("xxx"))))))
            .create();
    }
}
