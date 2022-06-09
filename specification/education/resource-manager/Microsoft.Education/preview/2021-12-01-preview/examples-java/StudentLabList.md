```java
import com.azure.core.util.Context;

/** Samples for StudentLabs ListAll. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentLabList.json
     */
    /**
     * Sample code: StudentLabList.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void studentLabList(com.azure.resourcemanager.education.EducationManager manager) {
        manager.studentLabs().listAll(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-education_1.0.0-beta.1/sdk/education/azure-resourcemanager-education/README.md) on how to add the SDK to your project and authenticate.
