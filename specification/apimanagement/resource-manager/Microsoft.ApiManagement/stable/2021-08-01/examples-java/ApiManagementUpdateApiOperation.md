Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.OperationContract;
import com.azure.resourcemanager.apimanagement.models.ParameterContract;
import com.azure.resourcemanager.apimanagement.models.RequestContract;
import com.azure.resourcemanager.apimanagement.models.ResponseContract;
import java.util.Arrays;

/** Samples for ApiOperation Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApiOperation.json
     */
    /**
     * Sample code: ApiManagementUpdateApiOperation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApiOperation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        OperationContract resource =
            manager
                .apiOperations()
                .getWithResponse("rg1", "apimService1", "echo-api", "operationId", Context.NONE)
                .getValue();
        resource
            .update()
            .withDisplayName("Retrieve resource")
            .withMethod("GET")
            .withUrlTemplate("/resource")
            .withTemplateParameters(Arrays.asList())
            .withRequest(
                new RequestContract()
                    .withQueryParameters(
                        Arrays
                            .asList(
                                new ParameterContract()
                                    .withName("param1")
                                    .withDescription(
                                        "A sample parameter that is required and has a default value of \"sample\".")
                                    .withType("string")
                                    .withDefaultValue("sample")
                                    .withRequired(true)
                                    .withValues(Arrays.asList("sample")))))
            .withResponses(
                Arrays
                    .asList(
                        new ResponseContract()
                            .withStatusCode(200)
                            .withDescription("Returned in all cases.")
                            .withRepresentations(Arrays.asList())
                            .withHeaders(Arrays.asList()),
                        new ResponseContract()
                            .withStatusCode(500)
                            .withDescription("Server Error.")
                            .withRepresentations(Arrays.asList())
                            .withHeaders(Arrays.asList())))
            .withIfMatch("*")
            .apply();
    }
}
```
