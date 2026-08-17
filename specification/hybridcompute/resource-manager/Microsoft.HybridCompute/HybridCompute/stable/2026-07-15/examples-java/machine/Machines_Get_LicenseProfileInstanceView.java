
import com.azure.resourcemanager.hybridcompute.models.InstanceViewTypes;

/**
 * Samples for Machines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/machine/Machines_Get_LicenseProfileInstanceView.json
     */
    /**
     * Sample code: Get Machine with License Profile Instance View.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        getMachineWithLicenseProfileInstanceView(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machines().getByResourceGroupWithResponse("myResourceGroup", "myMachine",
            InstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
