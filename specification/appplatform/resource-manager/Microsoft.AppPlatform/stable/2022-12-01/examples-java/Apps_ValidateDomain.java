import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.models.CustomDomainValidatePayload;

/** Samples for Apps ValidateDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Apps_ValidateDomain.json
     */
    /**
     * Sample code: Apps_ValidateDomain.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsValidateDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApps()
            .validateDomainWithResponse(
                "myResourceGroup",
                "myservice",
                "myapp",
                new CustomDomainValidatePayload().withName("mydomain.io"),
                Context.NONE);
    }
}
