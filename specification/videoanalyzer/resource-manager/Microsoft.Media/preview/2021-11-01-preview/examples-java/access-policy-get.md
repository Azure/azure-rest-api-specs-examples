Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AccessPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/access-policy-get.json
     */
    /**
     * Sample code: Gets an access policy entity.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getsAnAccessPolicyEntity(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.accessPolicies().getWithResponse("testrg", "testaccount2", "accessPolicyName1", Context.NONE);
    }
}
```
