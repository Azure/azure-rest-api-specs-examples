
import com.azure.resourcemanager.hybridcompute.fluent.models.SetupExtensionRequestInner;
import com.azure.resourcemanager.hybridcompute.models.MachineExtensionProperties;
import java.util.Arrays;

/**
 * Samples for ResourceProvider SetupExtensions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * extension/Extension_Add.json
     */
    /**
     * Sample code: Setup Machine Extensions.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void setupMachineExtensions(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.resourceProviders().setupExtensions("myResourceGroup", "myMachine",
            new SetupExtensionRequestInner().withExtensions(Arrays.asList(
                new MachineExtensionProperties().withPublisher("Microsoft.Azure.Monitoring")
                    .withType("AzureMonitorAgentLinux"),
                new MachineExtensionProperties().withPublisher("<extension_publisher>").withType("<extension_type>"))),
            com.azure.core.util.Context.NONE);
    }
}
