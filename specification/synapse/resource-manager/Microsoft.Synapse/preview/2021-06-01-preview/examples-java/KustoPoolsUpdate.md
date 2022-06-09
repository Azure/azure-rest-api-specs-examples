```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.AzureSku;
import com.azure.resourcemanager.synapse.models.KustoPool;
import com.azure.resourcemanager.synapse.models.SkuName;
import com.azure.resourcemanager.synapse.models.SkuSize;

/** Samples for KustoPools Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsUpdate.json
     */
    /**
     * Sample code: kustoPoolsUpdate.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsUpdate(com.azure.resourcemanager.synapse.SynapseManager manager) {
        KustoPool resource =
            manager
                .kustoPools()
                .getWithResponse("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", Context.NONE)
                .getValue();
        resource
            .update()
            .withSku(new AzureSku().withName(SkuName.STORAGE_OPTIMIZED).withCapacity(2).withSize(SkuSize.MEDIUM))
            .withEnableStreamingIngest(true)
            .withEnablePurge(true)
            .withWorkspaceUid("11111111-2222-3333-444444444444")
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
