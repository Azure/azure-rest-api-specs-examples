import com.azure.core.util.Context;

/** Samples for Profiles ListSupportedOptimizationTypes. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Profiles_ListSupportedOptimizationTypes.json
     */
    /**
     * Sample code: Profiles_ListSupportedOptimizationTypes.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilesListSupportedOptimizationTypes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .listSupportedOptimizationTypesWithResponse("RG", "profile1", Context.NONE);
    }
}
