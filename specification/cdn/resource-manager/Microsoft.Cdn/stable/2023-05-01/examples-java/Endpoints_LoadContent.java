import com.azure.resourcemanager.cdn.models.LoadParameters;
import java.util.Arrays;

/** Samples for Endpoints LoadContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Endpoints_LoadContent.json
     */
    /**
     * Sample code: Endpoints_LoadContent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsLoadContent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getEndpoints()
            .loadContent(
                "RG",
                "profile1",
                "endpoint1",
                new LoadParameters().withContentPaths(Arrays.asList("/folder1")),
                com.azure.core.util.Context.NONE);
    }
}
