
import com.azure.resourcemanager.cognitiveservices.models.ManagedComputeDeployment;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for ManagedComputeDeployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/UpdateVmManagedComputeDeployment.json
     */
    /**
     * Sample code: UpdateVmManagedComputeDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        updateVmManagedComputeDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        ManagedComputeDeployment resource = manager.managedComputeDeployments()
            .getWithResponse("resourceGroupName", "accountName", "gpt-oss-120b-byoc", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withSku(new Sku().withName("VmManagedCompute").withCapacity(2)).apply();
    }
}
