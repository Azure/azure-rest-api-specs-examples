import com.azure.core.util.Context;

/** Samples for Profiles ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Profiles_ListByResourceGroup.json
     */
    /**
     * Sample code: Profiles_ListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilesListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getProfiles().listByResourceGroup("RG", Context.NONE);
    }
}
