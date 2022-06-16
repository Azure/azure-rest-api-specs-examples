import com.azure.core.util.Context;

/** Samples for Profiles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Profiles_Delete.json
     */
    /**
     * Sample code: Profiles_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getProfiles().delete("RG", "profile1", Context.NONE);
    }
}
