
import com.azure.resourcemanager.cdn.models.AfdOriginUpdateParameters;
import com.azure.resourcemanager.cdn.models.EnabledState;

/**
 * Samples for AfdOrigins Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_Update.json
     */
    /**
     * Sample code: AFDOrigins_Update.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdOrigins().update("RG", "profile1", "origingroup1",
            "origin1", new AfdOriginUpdateParameters().withHostname("host1.blob.core.windows.net").withHttpPort(80)
                .withHttpsPort(443).withEnabledState(EnabledState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
