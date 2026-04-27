
/**
 * Samples for EdgeMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachines_Get.json
     */
    /**
     * Sample code: EdgeMachines_Get_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachinesGetMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachines().getByResourceGroupWithResponse("ArcInstance-rg", "machine-1",
            com.azure.core.util.Context.NONE);
    }
}
