
import com.azure.resourcemanager.cdn.models.AfdPurgeParameters;
import java.util.Arrays;

/**
 * Samples for AfdEndpoints PurgeContent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDEndpoints_PurgeContent.json
     */
    /**
     * Sample code: AFDEndpoints_PurgeContent.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDEndpointsPurgeContent(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdEndpoints().purgeContent(
            "RG", "profile1", "endpoint1", new AfdPurgeParameters().withContentPaths(Arrays.asList("/folder1"))
                .withDomains(Arrays.asList("endpoint1-abcdefghijklmnop.z01.azurefd.net")),
            com.azure.core.util.Context.NONE);
    }
}
