
import com.azure.resourcemanager.kusto.fluent.models.LanguageExtensionInner;
import com.azure.resourcemanager.kusto.models.AzureSku;
import com.azure.resourcemanager.kusto.models.AzureSkuName;
import com.azure.resourcemanager.kusto.models.AzureSkuTier;
import com.azure.resourcemanager.kusto.models.Identity;
import com.azure.resourcemanager.kusto.models.IdentityType;
import com.azure.resourcemanager.kusto.models.LanguageExtensionImageName;
import com.azure.resourcemanager.kusto.models.LanguageExtensionName;
import com.azure.resourcemanager.kusto.models.LanguageExtensionsList;
import com.azure.resourcemanager.kusto.models.PublicIpType;
import com.azure.resourcemanager.kusto.models.PublicNetworkAccess;
import java.util.Arrays;

/**
 * Samples for Clusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoClustersCreateOrUpdate
     * .json
     */
    /**
     * Sample code: KustoClustersCreateOrUpdate.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersCreateOrUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().define("kustoCluster").withRegion("westus").withExistingResourceGroup("kustorptest")
            .withSku(
                new AzureSku().withName(AzureSkuName.STANDARD_L16AS_V3).withCapacity(2).withTier(AzureSkuTier.STANDARD))
            .withIdentity(new Identity().withType(IdentityType.SYSTEM_ASSIGNED)).withEnableStreamingIngest(true)
            .withEnablePurge(true)
            .withLanguageExtensions(new LanguageExtensionsList().withValue(Arrays.asList(
                new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.PYTHON)
                    .withLanguageExtensionImageName(LanguageExtensionImageName.PYTHON3_10_8),
                new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.R)
                    .withLanguageExtensionImageName(LanguageExtensionImageName.R))))
            .withEnableDoubleEncryption(false).withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .withAllowedIpRangeList(Arrays.asList("0.0.0.0/0")).withEnableAutoStop(true)
            .withPublicIpType(PublicIpType.DUAL_STACK).create();
    }
}
