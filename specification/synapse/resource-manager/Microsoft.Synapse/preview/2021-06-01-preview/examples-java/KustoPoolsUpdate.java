
import com.azure.resourcemanager.synapse.models.AzureSku;
import com.azure.resourcemanager.synapse.models.KustoPool;
import com.azure.resourcemanager.synapse.models.SkuName;
import com.azure.resourcemanager.synapse.models.SkuSize;

/**
 * Samples for KustoPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsUpdate.
     * json
     */
    /**
     * Sample code: kustoPoolsUpdate.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsUpdate(com.azure.resourcemanager.synapse.SynapseManager manager) {
        KustoPool resource = manager.kustoPools().getWithResponse("synapseWorkspaceName", "kustoclusterrptest4",
            "kustorptest", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withSku(new AzureSku().withName(SkuName.STORAGE_OPTIMIZED).withCapacity(2).withSize(SkuSize.MEDIUM))
            .withEnableStreamingIngest(true).withEnablePurge(true).apply();
    }
}
