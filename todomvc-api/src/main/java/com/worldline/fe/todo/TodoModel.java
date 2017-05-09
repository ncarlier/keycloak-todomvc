/*
 * Copyright 2016 (C) Worldline - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */
package com.worldline.fe.todo;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.time.LocalDateTime;
import java.time.ZoneOffset;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.PrePersist;
import javax.persistence.PreUpdate;
import lombok.Data;

/**
 * Todo model.
 * @author Nicolas Carlier <nicolas.carlier@worldline.com>
 */
@Data
@Entity(name = "todo")
public class TodoModel {
  @Id
  @GeneratedValue(strategy = GenerationType.AUTO)
  private long id;
  
  @Column(nullable = false)
  private String title;
  
  @Column(nullable = false)
  private Boolean completed;
  
  @Column(columnDefinition = "TIMESTAMP DEFAULT CURRENT_TIMESTAMP")
  private LocalDateTime date;
  
  @Column(nullable = false)
  private String author;
  
  @PrePersist
  @PreUpdate
  protected void onCreateOrUpdate() {
    date = LocalDateTime.now();
  }
  
  @JsonProperty(value = "timestamp")
  public long getTimestamp() {
    return date.toEpochSecond(ZoneOffset.UTC);
  }
  
}
