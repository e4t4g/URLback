/*
 * UrlAPI
 * URL shortener REST API
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * InlineResponse200
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-06-05T22:46:22.672948+03:00[Europe/Moscow]")
public class InlineResponse200 {
  public static final String SERIALIZED_NAME_LINK = "link";
  @SerializedName(SERIALIZED_NAME_LINK)
  private String link;


  public InlineResponse200 link(String link) {
    
    this.link = link;
    return this;
  }

   /**
   * Get link
   * @return link
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "http://localhost:8006/32ebceda", value = "")

  public String getLink() {
    return link;
  }


  public void setLink(String link) {
    this.link = link;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse200 inlineResponse200 = (InlineResponse200) o;
    return Objects.equals(this.link, inlineResponse200.link);
  }

  @Override
  public int hashCode() {
    return Objects.hash(link);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse200 {\n");
    sb.append("    link: ").append(toIndentedString(link)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
