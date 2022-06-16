import com.azure.core.util.Context;

/** Samples for Profiles GenerateSsoUri. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Profiles_GenerateSsoUri.json
     */
    /**
     * Sample code: Profiles_GenerateSsoUri.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilesGenerateSsoUri(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .generateSsoUriWithResponse("RG", "profile1", Context.NONE);
    }
}
