import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.PurgeParameters;
import java.util.Arrays;

/** Samples for Endpoints PurgeContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_PurgeContent.json
     */
    /**
     * Sample code: Endpoints_PurgeContent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsPurgeContent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getEndpoints()
            .purgeContent(
                "RG",
                "profile1",
                "endpoint1",
                new PurgeParameters().withContentPaths(Arrays.asList("/folder1")),
                Context.NONE);
    }
}
