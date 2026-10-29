
/**
 * Samples for Profiles ListSupportedOptimizationTypes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_ListSupportedOptimizationTypes.json
     */
    /**
     * Sample code: Profiles_ListSupportedOptimizationTypes.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesListSupportedOptimizationTypes(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().listSupportedOptimizationTypesWithResponse("RG", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
