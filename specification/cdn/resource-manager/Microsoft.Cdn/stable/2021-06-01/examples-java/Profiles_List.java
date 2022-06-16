import com.azure.core.util.Context;

/** Samples for Profiles List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Profiles_List.json
     */
    /**
     * Sample code: Profiles_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getProfiles().list(Context.NONE);
    }
}
