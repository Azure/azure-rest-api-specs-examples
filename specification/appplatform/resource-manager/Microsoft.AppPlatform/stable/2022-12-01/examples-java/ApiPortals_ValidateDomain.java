import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.models.CustomDomainValidatePayload;

/** Samples for ApiPortals ValidateDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ApiPortals_ValidateDomain.json
     */
    /**
     * Sample code: ApiPortals_ValidateDomain.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void apiPortalsValidateDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApiPortals()
            .validateDomainWithResponse(
                "myResourceGroup",
                "myservice",
                "default",
                new CustomDomainValidatePayload().withName("mydomain.io"),
                Context.NONE);
    }
}
