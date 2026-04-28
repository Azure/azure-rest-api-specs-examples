
/**
 * Samples for DevicePools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/DevicePools_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DevicePools_Get_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void devicePoolsGetMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.devicePools().getByResourceGroupWithResponse("ArcInstance-rg", "fflisdaccdcoj",
            com.azure.core.util.Context.NONE);
    }
}
