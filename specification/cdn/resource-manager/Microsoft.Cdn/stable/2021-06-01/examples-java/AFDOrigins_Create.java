import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.fluent.models.AfdOriginInner;
import com.azure.resourcemanager.cdn.models.EnabledState;

/** Samples for AfdOrigins Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOrigins_Create.json
     */
    /**
     * Sample code: AFDOrigins_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdOrigins()
            .create(
                "RG",
                "profile1",
                "origingroup1",
                "origin1",
                new AfdOriginInner()
                    .withHostname("host1.blob.core.windows.net")
                    .withHttpPort(80)
                    .withHttpsPort(443)
                    .withOriginHostHeader("host1.foo.com")
                    .withEnabledState(EnabledState.ENABLED),
                Context.NONE);
    }
}
