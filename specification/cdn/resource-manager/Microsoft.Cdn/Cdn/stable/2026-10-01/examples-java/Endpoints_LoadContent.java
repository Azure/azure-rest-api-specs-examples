
import com.azure.resourcemanager.cdn.models.LoadParameters;
import java.util.Arrays;

/**
 * Samples for Endpoints LoadContent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Endpoints_LoadContent.json
     */
    /**
     * Sample code: Endpoints_LoadContent.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsLoadContent(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEndpoints().loadContent("RG", "profile1", "endpoint1",
            new LoadParameters().withContentPaths(Arrays.asList("/folder1")), com.azure.core.util.Context.NONE);
    }
}
