
import com.azure.resourcemanager.machinelearning.models.ManagedNetworkProvisionOptions;

/**
 * Samples for ManagedNetworkProvisions ProvisionManagedNetwork.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedNetwork/provision.json
     */
    /**
     * Sample code: Provision ManagedNetwork.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        provisionManagedNetwork(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.managedNetworkProvisions().provisionManagedNetwork("test-rg", "aml-workspace-name",
            new ManagedNetworkProvisionOptions().withIncludeSpark(false), com.azure.core.util.Context.NONE);
    }
}
