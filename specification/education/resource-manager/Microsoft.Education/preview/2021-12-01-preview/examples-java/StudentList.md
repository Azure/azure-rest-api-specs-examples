```java
import com.azure.core.util.Context;

/** Samples for Students List. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentList.json
     */
    /**
     * Sample code: StudentList.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void studentList(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .students()
            .list("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-education_1.0.0-beta.1/sdk/education/azure-resourcemanager-education/README.md) on how to add the SDK to your project and authenticate.
