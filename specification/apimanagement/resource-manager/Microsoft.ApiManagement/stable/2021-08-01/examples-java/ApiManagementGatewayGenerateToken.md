```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.GatewayTokenRequestContract;
import com.azure.resourcemanager.apimanagement.models.KeyType;
import java.time.OffsetDateTime;

/** Samples for Gateway GenerateToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayGenerateToken.json
     */
    /**
     * Sample code: ApiManagementGatewayGenerateToken.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGatewayGenerateToken(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gateways()
            .generateTokenWithResponse(
                "rg1",
                "apimService1",
                "gw1",
                new GatewayTokenRequestContract()
                    .withKeyType(KeyType.PRIMARY)
                    .withExpiry(OffsetDateTime.parse("2020-04-21T00:44:24.2845269Z")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
