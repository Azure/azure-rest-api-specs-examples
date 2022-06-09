```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.IdentityProviderContract;
import com.azure.resourcemanager.apimanagement.models.IdentityProviderType;

/** Samples for IdentityProvider Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateIdentityProvider.json
     */
    /**
     * Sample code: ApiManagementUpdateIdentityProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateIdentityProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        IdentityProviderContract resource =
            manager
                .identityProviders()
                .getWithResponse("rg1", "apimService1", IdentityProviderType.FACEBOOK, Context.NONE)
                .getValue();
        resource
            .update()
            .withClientId("updatedfacebookid")
            .withClientSecret("updatedfacebooksecret")
            .withIfMatch("*")
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
