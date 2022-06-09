```java
import com.azure.resourcemanager.kusto.models.AzureSku;
import com.azure.resourcemanager.kusto.models.AzureSkuName;
import com.azure.resourcemanager.kusto.models.AzureSkuTier;
import com.azure.resourcemanager.kusto.models.Identity;
import com.azure.resourcemanager.kusto.models.IdentityType;
import com.azure.resourcemanager.kusto.models.PublicIpType;
import com.azure.resourcemanager.kusto.models.PublicNetworkAccess;
import java.util.Arrays;

/** Samples for Clusters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClustersCreateOrUpdate.json
     */
    /**
     * Sample code: KustoClustersCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersCreateOrUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusters()
            .define("kustoCluster")
            .withRegion("westus")
            .withExistingResourceGroup("kustorptest")
            .withSku(new AzureSku().withName(AzureSkuName.STANDARD_L8S).withCapacity(2).withTier(AzureSkuTier.STANDARD))
            .withIdentity(new Identity().withType(IdentityType.SYSTEM_ASSIGNED))
            .withEnableStreamingIngest(true)
            .withEnablePurge(true)
            .withEnableDoubleEncryption(false)
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .withAllowedIpRangeList(Arrays.asList("0.0.0.0/0"))
            .withEnableAutoStop(true)
            .withPublicIpType(PublicIpType.DUAL_STACK)
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
